[![Build Status](https://travis-ci.org/mojlighetsministeriet/groups.svg?branch=master)](https://travis-ci.org/mojlighetsministeriet/groups)

[![Coverage Status](https://coveralls.io/repos/github/mojlighetsministeriet/groups/badge.svg?branch=master)](https://coveralls.io/github/mojlighetsministeriet/groups?branch=master)

# groups    

Handles the management of groups and their projects. This will let users collaborate by creating groups and from the groups create projects.

**NOTE: This service is still in it's experimental phase,** feel free to try it out, contribute pull requests but for now, expect the functionality to be somewhat limited. E.g. right now there is no way of creating the first administrator account yet. It will be added but we are not there yet. :)

## Philosofy     

* This service should do as few things as possible but do them well.
* This service should have as few customization options as humanly possible. There are so many well known alternatives out there that does these things but they require a week or more to get started with.
* This service should come with reasonable and as many default settings as possible to reduce configuration hell.
* This service is meant to run on Docker Swarm mode, it's okay to run it elsewhere but this is the main target.

## Docker image      

Our docker image is avaliable here https://hub.docker.com/r/mojlighetsministeriet/groups/.

## License

All the code in this project goes under the license GPLv3 (https://www.gnu.org/licenses/gpl-3.0.en.html).
