import { ReactElement, JSXElementConstructor, ReactFragment, ReactPortal } from 'react';
import './style.css';

function Button(props: { children?: string | number | boolean | ReactElement<any, string | JSXElementConstructor<any>> | ReactFragment | ReactPortal | null | undefined; }) {
    return (
      <button  className='btn-red'>{props.children}</button>
    );
  }

  export default Button; 