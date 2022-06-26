import userEvent from "@testing-library/user-event";
import React,{useState} from "react";
import './Css/OrderHistory.css'
import { getUserCookies } from "./Cookies";
import Products from "./Products";


function showOrders(orders){
    let i =0
    return orders.order.map((order)=>
    <div obj={order}  key={order.id_order}className="Orden">
        <ul className="ordenes">
            <li>        
                {order.det_orden.map((det_order)=>
                <div obj={det_order} key={i++} className="ProductoOrden">
                <ul className="Detalles">
                    <li> Productos:{det_order.product.name_product}           
                    <a className="precio_unitario">Precio Unitario:{det_order.price_det}</a>
                    <a className="cantidad">Cantidad:{det_order.cant_det}</a>
                    </li>
                </ul>
                </div>
                )
                }
                <a className="total">Total:{order.total}</a>
            </li>

        </ul>
    </div>)
    
}

export default function OrderHistory(){
    const[orders,setOrders]= useState([])
    const[token,setToken]= useState("")
    
    if(token==""){
        console.log("IF TOKEN")
        setToken(getUserCookies())
        console.log("USER IF",token)
    }
  
    const orderLoad=()=>{
        console.log("User",token)
        let url = 'http://localhost:3307/User/OrH/'+token
        if(token!=""&& token != undefined){
        fetch(url,{
            method: 'GET',
            headers: {'Content-Type':'application/IU'}
        })
        .then(response=>response.json())
        .then(response=>setOrders(response))
        }
    }
    if(orders==null||orders.length==0){
        orderLoad()
    }
    const handleSubmit= (event)=>{
        event.preventDefault()

    }
    return(
        <>
        <h3>OrderHistory</h3>
        <div onSubmit={handleSubmit}>
        </div>
            <div id="ordenes">
                {orders.order!=undefined ? showOrders(orders): <a>No OrderHistory For this User</a>}
            </div>
        
        </>
    )

}