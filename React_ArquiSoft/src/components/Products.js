import React, {useState,useEffect} from "react"
import addToCart from "./Cookies";


const loadOptions ={
    method:'GET',
    headers :{'Content-Type': 'application/'} 
}
async function getProducts(){
    return await fetch('http://127.0.0.1:3306/product/all',loadOptions)
    .then(response=>response.json())
}

function showProducts (products){
    return  products.map((product)=>
    <div obj={product} key={product.id_product} className="product">
        <ul className="Productos"> 
            <li>  {product.name_product}     
                <button onClick={()=> addToCart(product.id_product)}>  Agregar  </button> 
            </li> 
            <li>  Precio: ${product.cost}     </li>
            <li>   Categoria: {product.category}    </li>
            <li>   Stock: { product.stock}   </li>
            <li>Descripcion: {product.descripcion}</li>
        </ul>
     </div>)
}

export default function Products(){
    const[products,setProducts]= useState([]);
    const[keyPro,setKeyPro]=useState("");
    const[category,setCategory]=useState(false);
    const[name, setName] = useState("");
    const[cartProducts,setCartProducts]= useState ("");
    const loadOptions ={
        method:'GET',
        headers :{'Content-Type': 'application/'} 
    }
    const onSubmit= () =>{
        fetch('http://127.0.0.1:3306/product/all',loadOptions)
        .then(response=>response.json())
        .then(response=>setProducts(response));
    }
    if(products==null||products.length <=0){
        onSubmit()
    }
    
    const onChangeSearch = async(search) =>{
        setKeyPro(search.target.value);
    }

    const onChangeCat = async(category)=>{
        setCategory(category.target.checked);
    }
    
    const productsApi = async (keyPro) => {
        console.log(keyPro)
        if(keyPro==""){
            onSubmit();
            return
        }
        if(category==true){
        fetch('http://127.0.0.1:3306/product/cat/'+keyPro)
        .then(response=>response.json())
        .then(response=>{if(response==null){
            alert("Category not found")
            onSubmit()
            return
        }
        setProducts(response)
        })
        return
        }
      await fetch('http://127.0.0.1:3306/product/product/'+keyPro)
        .then(response=>response.json())
        .then(response=>{if(response==null){
            alert("Products not found")
            onSubmit()
            return
        }
        setProducts(response)
        })
        
       
    }   
    const handleSubmit = (event)=>{
        event.preventDefault();
    }
    return (
        <>
    <h3> PRODUCTOS </h3>
    <div onSubmit={handleSubmit}>
    <input type="text" id="search" placeholder="Search" onChange={onChangeSearch} value={keyPro} />
    <button type="submit" value="Buscar" onClick={()=>productsApi(keyPro)}>Buscar</button>
        <div>
            <span>busqueda por categorias</span>
            <input type="checkbox" id="category" onChange={onChangeCat} value={category} ></input>
        </div>    
    </div>
    <div id="product">
        {products.length>0 ? showProducts(products): <a>Nothing to show</a>}
     </div>
     </>
    )
    
}
