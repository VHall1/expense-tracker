FROM ruby:3.1.2-bullseye as base

WORKDIR /docker/app

RUN gem install bundler

COPY Gemfile* ./

RUN bundle install

ADD . /docker/app

CMD ["bundle","exec", "puma", "config.ru"]
