import React, { Component } from 'react';
import { Container } from 'flux/utils';
import { withRouter } from 'react-router-dom';
import Main from './Main';
import Navigation from './Navigation';

import * as appActions from '../actions/actions';
import appStore from '../stores/store';


class App extends Component {

    static getStores() {
        return [ appStore ];
    }

    static calculateState(prevState, props) {
        return appStore.getState();
    }

    render () {
        return <div>
            <Navigation { ...this.state }/>
            <Main
                { ...appActions }
                { ...this.state }
            />
        </div>;
    };
};

const component = new Container.create(App);
export default withRouter(component);

