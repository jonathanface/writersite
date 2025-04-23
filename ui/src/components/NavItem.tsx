import { useNavigate } from "react-router-dom";

interface NavItemProps {
    label: string;
    setActive: (value: string) => void;
    activeSection: string;
}

export const NavItem = (props: NavItemProps) => {
    const navigate = useNavigate();
    return (
        <h2>
            <button onClick={(event) => {
                event.preventDefault();
                props.setActive(props.label);
                navigate(props.label);
            }} className={props.activeSection === props.label ? 'sidebarActive' : ''}>{props.label.charAt(0).toUpperCase() + props.label.substring(1)}
            </button>
        </h2>
    );
};