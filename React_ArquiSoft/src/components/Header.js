import React, { useState } from 'react';
import { getUserCookies, logOutCookies } from "./Cookies";
import './Css/Header.css';


 function Header(){
    const [UserToken,setUserToken] = useState("")
    
    if(UserToken==""){
        setUserToken(getUserCookies())
    }
   
    return( 
    <nav>

        <div className = "header">
            <div>
            <img className="logo" src={require("./imagenes/mercado.png")} alt='logo'></img>
            </div>
        
            <div className="clases">
           <a className="botones" href="http://localhost:3000/products"  > PRODUCTOS </a>
           <a className="botones" href="#!"> CATEGORIAS </a>
           {UserToken!=undefined ? <a className="botones" href='#' onClick={()=>{logOutCookies(); window.location.reload()}}>LOG OUT</a>: <a className="botones" href="http://localhost:3000/"  > LOG IN </a>}
            </div>
            
            <div><a href="http://localhost:3000/cart">
            <img className="carrito" src={require("./imagenes/carrito.png")} alt='logocarrp'></img>
            </a>
            </div>
            {UserToken!=undefined? <a className='botones' href='http://localhost:3000/OrderHistory'>OrderHistory</a>: <></>}
            
        </div>
        

    </nav>

    )
}


export default Header