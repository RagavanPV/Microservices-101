from website.app import create_app
from python_logger import logger
import os


log = logger.get_instance('auth-server', "INFO", "2.0")
log.info('Application started')
log.info(os.environ)
mysql_uri = "mysql+pymysql://{$MYSQL_USER}:{$MYSQL_PASSWORD}@{$MYSQL_HOSTNAME}:{$MYSQL_PORT}/{$MYSQL_DB}"
mysql_uri = mysql_uri.replace("{$MYSQL_USER}", os.environ['MYSQL_USER'])
mysql_uri = mysql_uri.replace("{$MYSQL_PASSWORD}", os.environ['MYSQL_PASSWORD'])
mysql_uri = mysql_uri.replace("{$MYSQL_HOSTNAME}", os.environ['MYSQL_HOST'])
mysql_uri = mysql_uri.replace("{$MYSQL_PORT}", os.environ['E_COMMERCE_RELEASE_MYSQL_SERVICE_PORT'])
mysql_uri = mysql_uri.replace("{$MYSQL_DB}", os.environ['MYSQL_DB'])
app = create_app({
    'SECRET_KEY': 'secret',
    'OAUTH2_REFRESH_TOKEN_GENERATOR': True,
    'SQLALCHEMY_TRACK_MODIFICATIONS': False,
    'SQLALCHEMY_DATABASE_URI': mysql_uri,
    'SQLALCHEMY_ECHO': True
})
