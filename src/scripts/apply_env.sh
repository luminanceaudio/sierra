#!/bin/sh
set -e

replace_env_vars() {
    file_path="$1"

    rm -f "$file_path"
    cp "${file_path}.template" "$file_path"

    # Function to evaluate the environment variable value
    evaluate_env_var() {
        var_name="${1#@{{}"
        var_name="${var_name%\}\}}"
        eval "echo \$$var_name"
    }

    # Check if the file path is provided
    if [ -z "$file_path" ]; then
        echo "Error: File path not provided."
        return 1
    fi

    # Check if the file exists
    if [ ! -f "$file_path" ]; then
        echo "Error: File '$file_path' does not exist."
        return 1
    fi

    # Create a temporary file for writing the modified content
    temp_file="${file_path}.tmp"

    # Read the file line by line and write the modified content to the temporary file
    while IFS= read -r line; do
        # Search for @{{X}} pattern in the line
        while expr "$line" : '.*@{{[^}]*}}' > /dev/null; do
            # Extract the matched pattern and the variable name
            matched_string=$(expr "$line" : '.*\(@{{[^}]*}}\)')
            var_name=$(expr "$matched_string" : '@{{\([^}]*\)}}')
            # Evaluate the environment variable value
            evaluated_value=$(evaluate_env_var "$var_name")
            # Replace the matched pattern with the evaluated value
            line=$(echo "$line" | sed "s|$matched_string|$evaluated_value|")
        done
        # Append the modified line to the temporary file
        echo "$line" >> "$temp_file"
    done < "$file_path"

    # Move the temporary file to the original file path
    mv "$temp_file" "$file_path"
}

# Check if a file path is provided as a parameter
if [ -z "$1" ]; then
    echo "Error: File path not provided."
    exit 1
fi

replace_env_vars "$1"