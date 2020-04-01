FROM python3.8-alpine as base

FROM base as deps

WORKDIR /build
COPY requirements.txt /build
RUN pip install --install-option="--prefix=/install" -r requirements.txt

FROM base

WORKDIR /app
COPY --from=deps /install /usr/local
COPY app/* /app/

CMD ["python", "main.py"]

