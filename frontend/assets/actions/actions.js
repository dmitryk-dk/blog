import appDispatcher from '../utils/app-dispatcher';
import * as actionTypes from './action-types';
import * as http from '../utils/http';

export function initContainer (data) {
    appDispatcher.dispatch({
        actionType: actionTypes.ACTION_CONTAINER_INIT,
        data
    });
}


export function submit (post) {
    const url = `http://localhost:3030/post`;
    appDispatcher.dispatch({
        actionType: actionTypes.ACTION_SAVE_POST_REQUEST
    });
    http
        .post(url, post)
        .then(response => {
        response
            .json()
            .then(data => {
                if (data.success) {
                    appDispatcher.dispatch({
                        actionType: actionTypes.ACTION_SAVE_POST_SUCCESS,
                        post
                    });
                }
            })
            .catch(data => {
                console.log(data);
            })
        })
}

export function deletePost (id) {
    const url = `http://localhost:3030/delete`;
    appDispatcher.dispatch({
        actionType: actionTypes.ACTION_POST_DELETE_REQUEST,
        id
    });
    http
        .del(url, id)
        .then(response => {
            response
                .json()
                .then(data => {
                    if (data.success) {
                        appDispatcher.dispatch({
                            actionType: actionTypes.ACTION_POST_DELETE_SUCCESS,
                            id
                        });
                    }
                })
                .catch(data => {
                    console.log(data);
                })
        })
}
