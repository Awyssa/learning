# app.py
from flask import Flask, render_template

app = Flask(__name__)


@app.route('/')
def home():
    return render_template('base.html', content=render_template('home.html'))


@app.route('/about')
def about():
    return render_template('base.html', content=render_template('about.html'))


@app.route('/contact')
def contact():
    return render_template('base.html', content=render_template('contact.html'))


# Partial route handlers for HTMX requests
@app.route('/partial/about')
def partial_about():
    return render_template('about.html')


@app.route('/partial/contact')
def partial_contact():
    return render_template('contact.html')


if __name__ == '__main__':
    app.run(debug=True)