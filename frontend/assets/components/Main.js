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
                path={`/post:${props.id}`}
                children={ ( { match, ...rest } ) => extraProps.isSaving ?
                    <h1>Saving Post</h1>:
                    <Post { ...rest } { ...extraProps } /> }
            />
        </Switch>;

};


