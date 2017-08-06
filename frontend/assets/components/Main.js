import React from 'react';
import { Switch, Route } from 'react-router-dom';

import Home from './Home';
import Post from './Post';

export default (props) => {

    const extraProps = {
        submit: props.submit,
        ...props
    };
    return <Switch>
        <Route exact path="/" component={ Home } />
        <Route
            exact
            path={`/post:${props.id}`}
            render={ (props) => <Post{ ...props } { ...extraProps } /> }
        />
    </Switch>;
};

