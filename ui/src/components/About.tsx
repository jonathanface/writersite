import { useEffect, useState } from "react";

export const About = () => {

    const [content, setContent] = useState<string>("Loading");

    const fetchData = async () => {
        try {
            const response = await fetch('/html/about.html');
            if (!response.ok) {
                throw new Error('something went wrong: ' + response.body);
            }
            const payload = await response.text()
            setContent(payload);

        } catch (error: unknown) {
            console.error(error);
        }
    };

    useEffect(() => {
        fetchData();
    }, []);


    return (
        <div className="about sm:m-16 sm:max-w-[800px] about sm:p-8 pb-8 pt-4 text-base font-roboto overflow-hidden" dangerouslySetInnerHTML={{ __html: content }} />
    )
}