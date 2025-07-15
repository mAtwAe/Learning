FROM openjdk:24-jdk-slim
COPY . /usr/src/myapp
WORKDIR /the/workdir/path
RUN javac /usr/src/myapp/Main.java
CMD ["java", "-cp", "/usr/src/myapp", "Main"]