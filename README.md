<p align="center">
    <h1>CI-API-GO-4AL</h1>
</p>
<p align="center">
    <a href="https://github.com/ESGI-SCHOOL-4AL/ci-api-go-4al/actions"><img src="https://github.com/ESGI-SCHOOL-4AL/ci-api-go-4al/workflows/GoTest/badge.svg" alt="Test Status"/></a>
    <a href="https://github.com/ESGI-SCHOOL-4AL/ci-api-go-4al/actions"><img src="https://github.com/ESGI-SCHOOL-4AL/ci-api-go-4al/workflows/Deploy/badge.svg" alt="Deploy Status"/></a>
    <a href="https://goreportcard.com/report/github.com/ESGI-SCHOOL-4AL/ci-api-go-4al"><img src="https://goreportcard.com/badge/github.com/ESGI-SCHOOL-4AL/ci-api-go-4al" alt="Go Report"/></a>
    <a href='https://coveralls.io/github/ESGI-SCHOOL-4AL/ci-api-go-4al?branch=master'><img src='https://coveralls.io/repos/github/ESGI-SCHOOL-4AL/ci-api-go-4al/badge.svg?branch=master' alt='Coverage Status' /></a>
</p>

CI/CD project using Github Actions, based on a simple Web API made in Go.

- Automatic testing and lint on each commit and Pull Request.
- Automatic deployment on push on `master`.
- Coverage measurement.
- Repository permissions.
- Containerization of both development and production environments.

## Run locally

```
docker-compose -f docker-compose.dev.yml up
```

The API is accessible at `http://localhost:8080`.