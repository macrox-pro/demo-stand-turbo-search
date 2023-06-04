FROM python:3.9.6

RUN pip install rasa==3.5.9

RUN pip install natasha

WORKDIR /home/python/app
COPY . .

EXPOSE 50051

CMD ["python", "main.py"]