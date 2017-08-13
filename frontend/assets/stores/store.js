// utils
import { ReduceStore } from 'flux/utils';

// consts
import * as actionTypes from '../actions/action-types';
import appDispatcher from '../utils/app-dispatcher';


class Store extends ReduceStore {

    getInitialState () {
        return {
            //id:          null,
            //title:       '',
            //description: '',
            isSaving:    false,
            posts:       []
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
                    //title: action.post.title,
                    //description: action.post.description,
                    isSaving: false,
                    ...savedPosts.posts
                };
            case actionTypes.ACTION_POST_DELETE_SUCCESS:
                const updateByPosts = {
                    posts: [
                        ...state.posts.filter(post => post.id !== action.id)
                    ]
                };
                return {
                    ...state,
                    id: state.id - 1,
                    ...updateByPosts
                };
            default:
                return state;
        }
    }
}

export default new Store(appDispatcher);
