import React from 'react';
import { Link } from 'react-router-dom';

export default props =>
    <div className="div">
        <div className="top-bar" id="example-animated-menu" data-animate="hinge-in-from-top spin-out">
            <div className="top-bar-left">
                <ul className="dropdown menu" data-dropdown-menu>
                    <li className="menu-text">Go Example Blog</li>
                    <li><Link to='/'>Home</Link></li>
                    <li><Link to={`/post:${props.id}`}>Post</Link></li>
                </ul>
            </div>
        </div>
    </div>
