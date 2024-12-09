from flask import Flask, send_from_directory

# Create a Flask application instance
app = Flask(__name__)


# Serve the home page
@app.route("/")
def home():
    return send_from_directory(directory="static", path="index.html")


# Serve the about page
@app.route("/about")
def about():
    return send_from_directory(directory="static", path="about.html")


# Run the server
if __name__ == "__main__":
    app.run(debug=True)
