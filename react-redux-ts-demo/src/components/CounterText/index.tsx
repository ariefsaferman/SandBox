import React from 'react';

type CounterTextProps = {
  children?: string | number;
};

const CounterText: React.FC<CounterTextProps> = (props: CounterTextProps) => {
  const { children } = props;
  return <div>{children}</div>;
};

export default CounterText;
