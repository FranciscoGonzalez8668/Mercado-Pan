import React,{ useState} from "react"
import './Css/Login.css';
import { getUserCookies, loginCookies } from "./Cookies";

export default function Login(){
    const[emailOb,setEmail]= useState("");
    const[passwordOb,setPassword] = useState("");

    const onChangeEmail =  (email)=>{
        setEmail(email.target.value);
        
    }
    
    const onChangePas = (password)=>{
    setPassword(password.target.value)};

    
    const requestOptions={
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        
        body: JSON.stringify({email : emailOb, password : passwordOb })
    };

    const login = async()=>{
        fetch('http://127.0.0.1:3306/login',requestOptions)
        .then(response=>response.json())
        .then(response => {if (response.status == 400) {
            alert("user not found")
            window.location.reload();
        }else{
            loginCookies(response.token)
            console.log("Sacado de Cookies", getUserCookies())
            window.location.replace('/products') // mandar a pagina de productos
        }})
        
    };
   
    const handleSubmit= (event)=>{
        event.preventDefault();
        login();

    };

    return(
        <div className="login-form">
            <div className="form">
        <form className="login-form" onSubmit={handleSubmit}>
        <h1 className="Login">BIENVENIDO</h1>
        <h5 className="Login">Por Favor Introduzca su Usuario</h5>
        <div><input type="text" id="email" placeholder="email" onChange={onChangeEmail} value ={emailOb}></input></div>
        <div><input className="password" type="password" id="password" placeholder="password" onChange={onChangePas} value={passwordOb}></input></div>
        <button type="submit" value="Login" >Login</button>
        </form>
        </div>
        </div>
          );
    


}