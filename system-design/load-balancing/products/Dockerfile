FROM python:3.10-alpine
#Working Directory for the Flask application
WORKDIR /flask
#It will copy the Flask application files and the requirements.txt file
COPY . /flask
#It will install the dependencies for the Flask application
RUN pip install -r requirements.txt
#It starts the Flask application on Port 5000
CMD python app.py
