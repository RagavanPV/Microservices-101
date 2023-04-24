from website.app import create_app
from python_logger import logger


log = logger.get_instance('authlib', "INFO", "2.0")
log.info('Application started')
app = create_app({
    'SECRET_KEY': 'secret',
    'OAUTH2_REFRESH_TOKEN_GENERATOR': True,
    'SQLALCHEMY_TRACK_MODIFICATIONS': False,
    'SQLALCHEMY_DATABASE_URI': 'mysql+pymysql://root:root@host.docker.internal:33060/authDB',
})
