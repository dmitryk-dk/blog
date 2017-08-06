import appDispatcher from '../utils/app-dispatcher';
import * as actionTypes from './action-types';
import * as http from '../utils/http';

export function initContainer (data) {
    appDispatcher.dispatch({
        actionType: actionTypes.ACTION_CONTAINER_INIT,
        data
    });
}


export function submit (state) {
    const url = `http://localhost:3030/post`;
    appDispatcher.dispatch({
        actionType: actionTypes.ACTION_SAVE_POST_REQUEST
    });
    http
        .post(url, state)
        .then(response => {
        response
            .json()
            .then(data => {
                if (data.success) {
                    appDispatcher.dispatch({
                        actionType: actionTypes.ACTION_SAVE_POST_SUCCESS,
                        state
                    });
                }
            })
            .catch(data => {
                console.log(data);
            })
        })
}
