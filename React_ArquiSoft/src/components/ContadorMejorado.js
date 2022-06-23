import { type } from "@testing-library/user-event/dist/type";
import { useReducer } from "react"
import { contadorinit, contadorinitialState, contadorReducer } from "../reducers/contadorReducer";


const initialState = {contador: 0};
const init = (initialState) =>{

    return{
        contador: initialState.contador + 100,
    };

};



const TYPES = {

    INCREMENT: "INCREMENT",
    DECREMENT: "DECREMENT",
    RESET: "RESET"


}

function reducer(state,action){
    switch (action.type) {

        

        case TYPES.INCREMENT:
            return {contador: state.contador + 1}
            
    
        case TYPES.DECREMENT:
            return {contador: state.contador - 1}

        case TYPES.RESET:
            return initialState;
            
        default:
            return state;
    }

}



const ContadorMejorado = () =>{

    


    const[state, dispacht] = useReducer(contadorReducer,contadorinitialState, contadorinit);

    const sumar = () => dispacht({type:"INCREMENT"});
    const restar = () => dispacht({type:"DECREMENT"});
    const resetear = () => dispacht({type:"RESET"})
    return(
        <div style={{textAlight:"center"}}>
           <h2> Contador Mejorado Reducer</h2>    

            <nav>

                <button onClick={sumar}>  +  </button>
                <button onClick={resetear}>    0     </button>
                <button onClick={restar}>  -  </button>
                


            </nav>

            <h3>  {state.contador}  </h3>





     
        </div>
         

    )

}

export default ContadorMejorado