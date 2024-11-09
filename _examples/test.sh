# cd into _examples dir if not already there
if [[ $(pwd) != *"_examples" ]]; then
  cd _examples || exit
fi

# loop over each dir, but exclude "demo" dir
for dir in $(ls -d */ | grep -v demo); do
  echo "Running examples for '$dir'"
  cd "$dir" || exit

  # Loop over each dir
  for example in $(ls -d */); do
    echo "Running example for '$example'"
    cd "$example" || exit

    # Run the example
    go run main.go

    # Ask if the example was successful
    echo ""
    echo ""
    echo "Was the example successful? (y/n)"
    read -r success

    if [ "$success" != "y" ]; then
      echo ""
      echo "Exiting..."
      exit
    fi

    # Go back to the parent dir
    cd ..
  done

  cd ..
done
