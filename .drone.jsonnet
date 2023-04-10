local DockerBuid(name, push=false, trigger={}, when={}) = {
  kind: 'pipeline',
  type: 'kubernetes',
  name: name,
  steps: [
    {
      name: 'build' + (if push then ' and push' else ''),
      image: 'thegeeklab/drone-docker-buildx:23',
      privileged: true,
      settings: {
        registry: 'git.datti.to',
        repo: 'git.datti.to/dattito/timetable-dhbw',
        auto_tag: true,
        platforms: ['linux/amd64', 'linux/arm64', 'linux/arm/v7'],
        username: {
          from_secret: 'GIT_USER',
        },
        password: {
          from_secret: 'GIT_PASSWORD',
        },

      } + (
        if !push then {
          dry_run: true,
        } else {}
      ),
    } + (if when != {} then { when: when } else {}),
  ],
} + (if trigger != {} then { trigger: trigger } else {});

[
  DockerBuid('build', trigger={
    branch: {
      exclude: ['main'],
    },
    event: {
      exclude: ['tag'],
    },
  }),
  DockerBuid('build and push (tag)', true, trigger={
    event: ['tag'],
  }),
  DockerBuid('build and push (main-branch)', true, trigger={
    branch: ['main'],
  })
]
