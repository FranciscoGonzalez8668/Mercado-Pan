import React from "react";
import {Routes,Route, Router} from "react-router-dom"
import Login from "./Login";
import Products from "./Products";
import ShoppingCart from "./ShoppingCart";
import Home from "./Header"
import OrderHistory from "./OrderHistory";


export const Paginas =()=>{
    return(
        <section>
            <Routes>
                <Route path='/' exact element = {<Login/>}/>
                <Route path='/cart' exact element={<ShoppingCart/>}/>
                <Route path='/products' exact element={<Products/>}/>
                <Route path='/OrderHistory' exact element={<OrderHistory/>}></Route>
                
            </Routes>
        </section>


    )
}
