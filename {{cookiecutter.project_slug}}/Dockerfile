FROM {{cookiecutter.docker_build_image}} AS build

# no Debian packages
RUN curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | bash -s -- --to /usr/local/bin

ADD . /go/src/github.com/{{cookiecutter.github_username}}/{{cookiecutter.project_slug}}
WORKDIR /go/src/github.com/{{cookiecutter.github_username}}/{{cookiecutter.project_slug}}

RUN just install-static


FROM {{cookiecutter.docker_run_image}}

LABEL APP="{{cookiecutter.project_slug}}"
LABEL REPO="https://github.com/{{cookiecutter.github_username}}/{{cookiecutter.project_slug}}"

COPY --from=build /go/bin/{{cookiecutter.project_slug}} /
CMD ["/{{cookiecutter.project_slug}}"]
