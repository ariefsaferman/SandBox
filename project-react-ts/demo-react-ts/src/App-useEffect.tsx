import React, { useEffect, useState } from 'react';


import './App.css';

function App() {    
    // console.log('Inside App()')
    const [num, setNum] = useState(0)


    useEffect (() => {
        // console.warn('inside useEffect')
        
        const timer = window.setInterval(() => {
            //because of something called "closure"
            // num = 0
            // setNum(num + 1);
            setNum((prevState) => prevState + 1); 
        }, 1000);

        //need to clear the interval for clean up
        return () => window.clearInterval(timer); 
    }, []);

    // console.log('Inside app() - 2')
    return (
        <div className='App'>
            APP
            <p>{num}</p>
            <button onClick={() => setNum(num + 1)}>btn</button>
        </div>
    );
}

export default App;
