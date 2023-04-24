FROM python:3.8-slim-buster

WORKDIR /app

COPY requirements.txt requirements.txt
COPY ./dependencies/python_logger-1.0-py3-none-any.whl python_logger-1.0-py3-none-any.whl
RUN pip3 install -r requirements.txt
RUN pip3 install python_logger-1.0-py3-none-any.whl

COPY . .
EXPOSE 5000
CMD [ "python3", "-m" , "flask", "run", "--host=0.0.0.0"]