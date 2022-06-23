import { TYPES } from "../actions/contadorActions";

export const contadorinitialState = {contador: 0};
export const contadorinit = (initialState) =>{

    return{
        contador: initialState.contador + 100,
    };

};



export function contadorReducer(state,action){
    switch (action.type) {

        

        case TYPES.INCREMENT:
            return {contador: state.contador + 1}
            
    
        case TYPES.DECREMENT:
            return {contador: state.contador - 1}

        case TYPES.RESET:
            return contadorinitialState;
            
        default:
            return state;
    }

}