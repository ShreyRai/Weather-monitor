Important points:
Attack Surface: "I used Alpine and Multi-stage builds to keep the image under 20MB and reduce the number of installed packages that could have vulnerabilities."

The Root Issue: "I implemented a non-root user in the Backend Dockerfile to adhere to the Principle of Least Privilege."

Layer Caching: "I structured my COPY commands to ensure that source code changes don't trigger a full re-download of dependencies."

Port Mapping vs. Networking: "I used -p to expose the UI to the user, but the containers communicate internally via the weather-net bridge using Docker's internal DNS."
