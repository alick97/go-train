import logging
from flask import Flask, request

app = Flask(__name__)

logging.basicConfig(
    level=logging.DEBUG,  # Set the overall logging level
    format='%(asctime)s - %(levelname)s - %(message)s'
)

@app.route("/")
def hello():
    v = request.args.get('v', '')
    app.logger.info(f"v: {v}")
    return f"v: {v}"

if __name__ == "__main__":
    app.run(debug=False, host='0.0.0.0', port=8000) #For testing only.  Don't use debug=True in production.
