import pandas as pd
import plotly.graph_objects as go
import plotly.subplots as sp
import bars.ohlc as ohlc


def main():
    filepath = "/home/grahamtonysmith/code/go-alg-zigzag/data/gc.csv"
    
    header = 0
    names=[
        "date",
        "time",
        "open",
        "high",
        "low",
        "close",
        "volume",
    ]
    
    df = pd.read_csv(filepath_or_buffer=filepath, header=header, names=names)
    
    ohlc = go.Ohlc(x=df.index, open=df["open"], high=df["high"], low=df["low"], close=df["close"], showlegend=False)
    volume = go.Bar(x=df.index, y=df["volume"], showlegend=False, marker={"color": "rgba(128,128,128,0.5)"})

    fig = sp.make_subplots(specs=[[{"secondary_y": True}]])

    fig.add_trace(ohlc, secondary_y=True)
    fig.add_trace(volume, secondary_y=False, )
    fig.update_layout(title="GC", xaxis={"rangeslider": {"visible": False}})
    fig.update_yaxes(title="Price/$", secondary_y=True, showgrid=True)
    fig.update_yaxes(title="Volume/Contracts", secondary_y=False, showgrid=False)
    
    fig.show()
    
if __name__ == "__main__":
    main()
