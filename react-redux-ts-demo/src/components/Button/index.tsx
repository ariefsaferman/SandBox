import React from 'react';

type ButtonProps = {
  children?: string | number;
  onClickHandler?: () => void;
};

const Button: React.FC<ButtonProps> = (props: ButtonProps) => {
  const { children, onClickHandler } = props;
  return <button onClick={onClickHandler}>{children}</button>;
};

export default Button;
