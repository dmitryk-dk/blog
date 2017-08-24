import React, { Component } from 'react';

export default class PostForm extends Component {

    constructor (props) {
        super (props);
        console.log(props);
        this.state = {
            id:          props.id,
            title:       props.title,
            description: props.description
        }
    }

    render () {
        console.log(this.state);
        const { id, title, description } = this.state;
        return <form className="text-center" onSubmit={ (event) => this._handleSubmit(event) }>
            <div className="row">
                <div className="medium-12 columns">
                    <span className="warning label">{ id }</span>
                </div>
                <div className="medium-12 columns">
                    <label>Title
                        <input
                            type="text"
                            placeholder="Please add title..."
                            name="title"
                            value={ title }
                            onChange={ ({ target }) => this._handleChange(target.name, target.value) }
                        />
                    </label>
                </div>
                <div className="medium-12 columns">
                    <label>Content
                        <textarea
                            type="text"
                            name="description"
                            placeholder="Content text..."
                            className="post-textarea"
                            value={ description }
                            onChange={ ({ target }) => this._handleChange(target.name, target.value) }
                        />
                    </label>
                </div>
                <button
                    className="button success round"
                >
                    Submit
                </button>
            </div>
        </form>;
    };

     _handleChange = (name, value) => {
        this.setState(state => ({
            ...state,
            [name]: value
        }));
    };

    _handleSubmit = (event) => {
        const { submit } = this.props;
        event.preventDefault();
        submit(this.state);
    };
}

