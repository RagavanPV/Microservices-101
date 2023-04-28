import os
from flask import Flask
from gevent.pywsgi import WSGIServer
from .models import db
from .oauth2 import config_oauth
from .routes import bp
from urllib.request import urlopen
from python_logger import logger

log = logger.get_instance('auth-server', "INFO", "2.0")

def create_app(config=None):
    app = Flask(__name__)
    os.environ['AUTHLIB_INSECURE_TRANSPORT'] = '1'

    # load default configuration
    app.config.from_object('website.settings')

    # load environment configuration
    if 'WEBSITE_CONF' in os.environ:
        app.config.from_envvar('WEBSITE_CONF')

    # load app specified configuration
    if config is not None:
        if isinstance(config, dict):
            app.config.update(config)
        elif config.endswith('.py'):
            app.config.from_pyfile(config)
    setup_app(app)
    # Debug/Development
    app.run(debug=True, host="0.0.0.0", port="5000")
    # Production
    # http_server = WSGIServer(('127.0.0.1', 5000), app)
    # http_server.serve_forever()
    return app


def setup_app(app):
    # Create tables if they do not exist already
    log.info('Create tables if they do not exist already')
    @app.before_first_request
    def create_tables():
        db.create_all()
    log.info('Create tables done')
    db.init_app(app)
    log.info('init app done')
    config_oauth(app)
    log.info('configure routes')
    app.register_blueprint(bp, url_prefix='')
