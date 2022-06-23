import { TYPES } from "../actions/shoppingActions"

export const shoppingInitialState = {

        products: [ //DE ACA PUEDO TRAER LOS PRODUCTOS CON UN API
        //LOS HAGO YO PARA PROBAR
        
         { id_product: 1, name_product: "Producto 1", cost:100, stock: 5, category: "electronica",descripcion:"lindo" },
         { id_product: 2, name_product: "Producto 2", cost:100, stock: 5, category: "electronica",descripcion:"lindo" },
         { id_product: 3, name_product: "Producto 6", cost:1020, stock: 5, category: "electronica",descripcion:"lindo" },
         { id_product: 4, name_product: "Producto 9", cost:10, stock: 5, category: "electronica",descripcion:"lindo" },
         { id_product: 5, name_product: "Producto 4", cost:200, stock: 5, category: "electronica",descripcion:"lindo" },

        ],
        cart:[],

};


export function shoppingReducer(state,action){

        switch (action.type) {
                case TYPES.ADD_TO_CART :{
                        let newItem = state.products.find( (product) => product.id_product === action.payload);
                       console.log(newItem);
                       let itemInCart = state.cart.find((item) => item.id_product === newItem.id_product);//aca para q no repita muchas veces un producto


                       return itemInCart
                       ? {
                           ...state,
                           cart: state.cart.map((item) =>
                             item.id_product === newItem.id_product
                               ? { ...item, quantity: item.quantity + 1 }
                               : item
                           ),
                         }
                       : {
                           ...state,
                           cart: [...state.cart, { ...newItem, quantity: 1 }],
                         };

                }
                
                case TYPES.REMOVE_ONE_FROM_CART: {


                        let itemToDelete = state.cart.find((item) => item.id_product === action.payload);

                         return itemToDelete.quantity > 1
                                   ? {
                                      ...state,
                                               cart: state.cart.map((item) =>
                                 item.id_product === action.payload
                                  ? { ...item, quantity: item.quantity - 1 }
                                             : item
                                                         ),
                                }
                                               : {
                                                    ...state,
                                           cart: state.cart.filter((item) => item.id_product !== action.payload),
                                                          };


                }


                case TYPES.REMOVE_ALL_FROM_CART: {

                        return {
                                ...state,
                                cart: state.cart.filter((item) => item.id_product !== action.payload),
                              };














                }
                
                case TYPES.CLEAR_CART: 
                        return shoppingInitialState;
                
                default:
                         return state;
        }  
}