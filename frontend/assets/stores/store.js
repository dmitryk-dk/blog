// utils
import { ReduceStore } from 'flux/utils';

// consts
import * as actionTypes from '../actions/action-types';
import appDispatcher from '../utils/app-dispatcher';


class Store extends ReduceStore {

    getInitialState () {
        return {
            isSaving:    false,
            posts:       [],
            id:          0
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
                    posts: action.data,
                    id:    action.data.length
                };
            case actionTypes.ACTION_SAVE_POST_REQUEST:
                return {
                    ...state,
                    isSaving: true
                };
            case actionTypes.ACTION_SAVE_POST_SUCCESS:
                const savedPosts = {
                    ...state,
                    ...state.posts.push(action.post)
                };
                return {
                    ...state,
                    id: state.id + 1,
                    isSaving: false,
                    ...savedPosts.posts
                };
            case actionTypes.ACTION_POST_DELETE_REQUEST:
                return {
                    ...state,
                    isSaving: true,
                };
            case actionTypes.ACTION_POST_DELETE_SUCCESS:
                const updateByPosts = {
                    posts: [
                        ...state.posts.filter(post => post.id !== action.id)
                    ],
                    id: state.id - 1,
                    isSaving: false,
                };
                return {
                    ...state,
                    ...updateByPosts
                };
            default:
                return state;
        }
    }
}

export default new Store(appDispatcher);
