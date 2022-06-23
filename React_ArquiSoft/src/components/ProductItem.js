import React from  "react"
export const ProductItem = (
    {
        id_product,
        name_product,
        price,
        stock,
        category,
        descripcion,
    }
    ) => {
    
    const addToCart = ()=>{
        console.log('coockies')
    }
    
    
    return <div style={{border:"thin solid gray", padding:"1rem"}}> 
    
    <h4>  {name_product}  </h4> 
    <h5>   ${price}.00      </h5>

    <button onClick={()=> addToCart(id_product)}>  Agregar  </button>
     </div>


    


};


export default ProductItem;