image=yaml-testing-cpp-yamlcpp
name=$(basename $1)

[ -z "$BUILD" ] || docker build -t $image .

cd ../../
docker run -v $PWD/testdata/$name:/$name $image /yaml-cpp/main /$name
