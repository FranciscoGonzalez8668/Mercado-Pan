import logo from './logo.svg';
import './App.css';
import {Paginas} from './components/Paginas'

import {BrowserRouter as Router} from 'react-router-dom';
import Products from './components/Products'
import Contador from './components/Contador';
import ContadorMejorado from './components/ContadorMejorado';
import ShoppingCart from './components/ShoppingCart';
import Header from './components/Header';


function App() {

  return (
    <div className='App'>
     <Header/>
      <Router>
        <Paginas/>
      </Router> 
      

     </div>
  );
}

export default App;

/*import React from 'react';
import*/
/* <hr/>   <ShoppingCart></ShoppingCart>
    <hr/>   <ContadorMejorado></ContadorMejorado>   
    <hr/>   <Contador></Contador> */




    //---------------------LO DEL AAGUS----------------------------
    /* return (
    <div>
    <h1> useReducer</h1>
    
    <Products/>
    
     
     </div>
  );
}*/