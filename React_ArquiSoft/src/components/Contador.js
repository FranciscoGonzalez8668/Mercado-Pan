import { type } from "@testing-library/user-event/dist/type";
import { useReducer } from "react"

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



const Contador = () =>{

    

  //  const[contador, setContador] = useState(0)
    const[state, dispacht] = useReducer(reducer,initialState, init);

   // const sumar = () => setContador(contador  + 1)
    const sumar = () => dispacht({type:"INCREMENT"});
    const restar = () => dispacht({type:"DECREMENT"});
    const resetear = () => dispacht({type:"RESET"})
    return(
        <div style={{textAlight:"center"}}>
           <h2> Contador Reducer</h2>    

            <nav>

                <button onClick={sumar}>  +  </button>
                <button onClick={resetear}>    0     </button>
                <button onClick={restar}>  -  </button>
                


            </nav>

            <h3>  {state.contador}  </h3>





     
        </div>
         

    )

}

export default Contador