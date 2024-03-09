[ -z "$BUILD" ] || gradle assemble
java -jar ./app/build/libs/app.jar $1
