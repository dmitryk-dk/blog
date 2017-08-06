// utils
import { ReduceStore } from 'flux/utils';

// consts
import * as actionTypes from '../actions/action-types';
import appDispatcher from '../utils/app-dispatcher';


class Store extends ReduceStore {

    getInitialState () {
        return {
            id:          null,
            title:       '',
            description: ''
        };
    }

    /**
     * @param {Object} state
     * @param {Object} action
     * @returns {Object}
     */
    reduce (state, action) {
        switch (action.actionType) {
            case actionTypes.ACTION_CONTAINER_INIT:
                return {
                    ...state,
                    ...action.data
                };
            case actionTypes.ACTION_SAVE_POST_REQUEST:
                return {
                    ...state,
                };
            case actionTypes.ACTION_SAVE_POST_SUCCESS:
                console.log(action)
                return {
                    ...state,
                    id: state.id + 1,
                    title: action.state.title,
                    description: action.state.description,
                };
            default:
                return state;
        }
    }
}

export default new Store(appDispatcher);
