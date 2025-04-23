import { JSX, useCallback, useEffect, useState } from "react";
import { Post } from "../types/Post";
import { NewsEntry } from "./NewsEntry";

export const News = () => {

    const [content, setContent] = useState<JSX.Element[] | string>("Loading...");

    const processNewsEntries = async (jsonData: Post[]) => {
        const content = [];
        let idx = 0;
        for (const data of jsonData) {
            const entryFile = `html/news/${data.file}`;
            try {
                const response = await fetch(entryFile);
                const html = await response.text();
                console.log(`Content of ${entryFile}:\n`, html);
                const postedDate = new Date(data.posted_on * 1000);
                content.push(<NewsEntry id={idx} htmlString={html} title={data.title} posted_on={postedDate} />)
                idx++;
            } catch (error) {
                console.error(`Error fetching ${entryFile}:`, error);
            }
        }

        setContent(content);
    }

    const fetchData = useCallback(async () => {
        try {
            const response = await fetch('/api/news');
            if (!response.ok) {
                throw new Error('something went wrong: ' + response.body);
            }
            const payload = await response.json() as Post[];
            processNewsEntries(payload)

        } catch (error: unknown) {
            console.error(error);
        }
    }, []);

    useEffect(() => {
        fetchData();
    }, [fetchData]);


    return (
        <div className={"news"}>{content}</div>
    )
}