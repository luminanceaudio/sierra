import React, { useEffect } from 'react';
import Input from '../Input/Input';

export type SearchProps = {
  value: string;
  onChange: (value: string) => void;
};

function Search({
  value: bufferedValue,
  onChange: onChangeBuffered,
}: SearchProps): React.ReactNode {
  const [value, setValue] = React.useState('');

  // Call onChange but with a buffer of 500 ms
  useEffect(() => {
    const intervalId = setInterval(() => {
      onChangeBuffered(value);
    }, 90);

    // Cleanup
    return () => clearInterval(intervalId);
  }, [onChangeBuffered, value]);

  // Set value if changed from outside
  useEffect(() => {
    setValue(bufferedValue);
  }, [bufferedValue]);

  return (
    <Input
      value={value}
      onChange={setValue}
      placeholder="Search samples"
      style={{
        paddingTop: 15,
        paddingBottom: 15,
        paddingLeft: 20,
        borderRadius: 50,
      }}
    />
  );
}

export default Search;
