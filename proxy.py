from flask import Flask, jsonify, request, Response
import threading
import time
import requests
import json
from dateutil import parser

app = Flask(__name__)

latest_release = {"release_ts": 0, "release_version": ""}
active_downloads = {}
total_downloaded_bytes = 0  # All download files

@app.route('/status', methods=['GET'])
def status_page():
    global active_downloads, total_downloaded_bytes

    timeout = 10
    current_time = time.time()
    active_downloads_filtered = {ip: ts for ip, ts in active_downloads.items() if current_time - ts < timeout}
    active_downloads = active_downloads_filtered

    # Calculate data downloaded shitt
    if total_downloaded_bytes >= 1_000_000_000:
        total_downloaded = f"{total_downloaded_bytes / 1_000_000_000:.2f} GB"
    else:
        total_downloaded = f"{total_downloaded_bytes / 1_000_000:.2f} MB"

    # JSON
    return jsonify({
        "latest_release_version": latest_release.get("release_version", "N/A"),
        "proxy_version": "1.6.2",
        "active_downloads": len(active_downloads),
        "total_downloaded": total_downloaded
    })

@app.route('/', methods=['GET'])
def handle_forward_request():
    global active_downloads, total_downloaded_bytes

    target_url = request.args.get('url')
    if target_url:
        ip = request.remote_addr
        active_downloads[ip] = time.time()

        try:
            with requests.get(target_url, stream=True, timeout=(5, 10)) as response:
                response.raise_for_status()
                content = response.content
                downloaded_size = len(content)  # download size
                total_downloaded_bytes += downloaded_size  # Update all 
                return Response(
                    content,
                    content_type=response.headers.get('Content-Type', 'application/octet-stream'),
                    status=response.status_code
                )
        except requests.RequestException as e:
            return Response(f"Error fetching the requested URL: {e}", status=500)
    return Response("Unsupported request", status=400)

@app.route('/client/builds.json', methods=['GET'])
def handle_builds_json():
    global latest_release
    fetch_latest_release()  # Fetch the version from github
    return jsonify(latest_release)

def fetch_latest_release():
    global latest_release
    try:
        with open('release.json', 'r') as f:
            data = json.load(f)
            release_time_str = data.get("release_time")
            release_version = data.get("release_version", "N/A")

            if release_time_str:
                release_ts = parser.isoparse(release_time_str).timestamp()
            else:
                release_ts = time.time()

            latest_release = {
                "release_ts": release_ts,
                "release_version": release_version
            }
    except (FileNotFoundError, json.JSONDecodeError, ValueError) as e:
        print(f"Error reading release.json: {e}")
        latest_release = {
            "release_ts": time.time(),
            "release_version": "N/A"
        }

if __name__ == '__main__':
    port = 5090
    fetch_latest_release() 
    app.run(host='0.0.0.0', port=port)