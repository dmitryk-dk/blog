import React from 'react';
import { Switch, Route } from 'react-router-dom';

import Home from './Home';
import PostForm from './PostForm';

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
                    <div>
                        <PostForm { ...rest } { ...extraProps } />
                        {
                            props.posts.map((post, i) => {
                                return (
                                    <div className="callout secondary" key={ i }>
                                        <div className="row">
                                            <div className="small-12 columns">
                                                <span className="success badge">{ post.id }</span>
                                                <h5>{ post.title }</h5>
                                                <div>{ post.description }</div>
                                            </div>
                                            <div className="small-12 columns text-right">
                                                <button
                                                    className="button large alert"
                                                    onClick={ () => props.deletePost(post.id) }>
                                                    Delete
                                                </button>
                                            </div>
                                        </div>
                                    </div>
                                );
                            })
                        }
                    </div>
                 }
            />
        </Switch>;

};


