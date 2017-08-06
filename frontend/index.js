import React from 'react';
import { render } from 'react-dom';
import { BrowserRouter } from 'react-router-dom';
import App from './assets/components/App';
import * as appActions from './assets/actions/actions';
import './assets/scss/index.scss';

function initApp () {
    const containerEl = document.getElementById('app');
    const data = containerEl.getAttribute('data-react');
    if (data) {
        try {
            return JSON.parse(data) || {};
        } catch (exception) {
            console.error('JSON parse error: ' + exception);
        }
    }
    return {};
}


(function init () {
    appActions.initContainer(initApp());
    render(
    <BrowserRouter>
        <App />
    </BrowserRouter>,
    document.getElementById('app'))
})();
