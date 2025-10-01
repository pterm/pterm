#!/usr/bin/env bash

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
    
    # Read single character without waiting for return
    read -n 1 -r success
    echo  # Add newline after single character input

    if [ "$success" != "y" ]; then
      echo ""
      # Output red text with full path to failed example
      echo -e "\033[31mFailed example: $(pwd)\033[0m"
      echo ""
      echo "Exiting..."
      exit
    fi

    # Go back to the parent dir
    cd ..
  done

  cd ..
done
