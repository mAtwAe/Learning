from fastapi import FastAPI, Response
from feedgenerator import Rss201rev2Feed
from datetime import datetime

app = FastAPI()

@app.get("/rss.xml")
def get_rss():
    # Create RSS feed
    feed = Rss201rev2Feed(
        title="Example Blog",
        link="https://www.example.com",
        description="This is a sample RSS feed for a blog",
        language="en",
        lastBuildDate=datetime.now()
    )

    # Add items
    feed.add_item(
        title="How to Cook Perfect Rice",
        link="https://www.example.com/blog/perfect-rice",
        description="A guide to cooking rice perfectly every time.",
        pubdate=datetime(2025, 6, 18, 10, 0),
        unique_id="https://www.example.com/blog/perfect-rice"
    )

    feed.add_item(
        title="5 Java Tips for Beginners",
        link="https://www.example.com/blog/java-tips",
        description="Tips to help Java beginners write better code.",
        pubdate=datetime(2025, 6, 17, 9, 0),
        unique_id="https://www.example.com/blog/java-tips"
    )

    # Generate XML
    from io import StringIO
    output = StringIO()
    feed.write(output, 'utf-8')
    xml_content = output.getvalue()

    return Response(content=xml_content, media_type="application/rss+xml")
