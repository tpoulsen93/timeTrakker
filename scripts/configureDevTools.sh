#!/bin/zsh

rulem() {
  if [ $# -eq 0 ]; then
    echo "Usage: rulem <message> [<rule character>]"
    return 1
  fi

  # fill line with ruler characters ($2, default "-"), reset cursor, move 2 cols right, print message
  printf -v _hr "%*s" $(tput cols) &&
    echo "" &&
    echo -en "${_hr// /${2--}}" &&
    echo -e "\r\033[2C$1" &&
    echo ""
}

didSucceed() {
  if (( $? == 0 )); then
    echo "SUCCESS"
  else
    echo "FAIL"
  fi
}

rulem "[ setting up postgresql ]" =

sudo service postgresql start
didSucceed

echo "Creating user..."
sudo runuser -l postgres -c "createuser -w -d -s $(whoami)"
didSucceed

echo "Creating database..."
createdb $(whoami)
didSucceed

echo "Creating password..."
psql -c "alter user $(whoami) with password 'secret';"
didSucceed

echo "Creating tables..."
psql -f db/schema.sql
