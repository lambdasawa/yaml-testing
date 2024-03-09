FROM ubuntu:latest

RUN apt-get update && apt-get install -y \
  g++ \
  cmake \
  git

RUN git clone https://github.com/jbeder/yaml-cpp.git

WORKDIR /yaml-cpp
RUN mkdir build \
  && cd build \
  && cmake .. \
  && make \
  && make install

COPY ./main.cpp .

RUN g++ -std=c++14 main.cpp -lyaml-cpp -o main

CMD ["./main"]
