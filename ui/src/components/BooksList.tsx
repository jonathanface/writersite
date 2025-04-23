import { JSX, useCallback, useEffect, useState } from "react";
import { Book } from "../types/Book";

export const BooksList = () => {

    const formatListItems = (data: Book[]) => {
        const response = data.map((item: Book) => {
            return (
                <div className="sm:m-16 book">
                    <div className="text-center sm:max-w-[600px]">
                        <h2 className="text-[1.5rem] text-left m-auto sm:max-w-[400px]"><a href={item.link}>{item.title}</a></h2>
                        <div className="text-left m-auto sm:max-w-[400px]">{new Date(item.released_on * 1000).toLocaleString('default', { month: 'long', year: 'numeric' })}</div>
                        <div className="sm:max-w-[400px] m-auto"><a href={item.link}><img alt={item.title} src={'img/covers/' + item.image} /></a></div>
                    </div>
                    <div className="p-8 pl-0 pr-0 pt-4 text-base font-roboto sm:max-w-[600px] text-justify">{item.description}</div>
                    <hr />
                </div>);
        });
        return (<div>{response}</div>)
    };

    const [content, setContent] = useState<JSX.Element | string>("Loading...");

    const fetchData = useCallback(async () => {
        try {
            const response = await fetch('/api/books');
            if (!response.ok) {
                throw new Error('something went wrong: ' + response.body);
            }
            const payload = await response.json();
            if (Array.isArray(payload)) {
                setContent(formatListItems(payload));
            } else {
                setContent(<div className="sm:m-16 text-base font-roboto sm:max-w-[600px]">{payload}</div>);
            }

        } catch (error: unknown) {
            console.error(error);
        }
    }, []);

    useEffect(() => {
        fetchData();
    }, [fetchData]);


    return (
        <div className="books">{content}</div>
    )
}