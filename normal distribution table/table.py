import numpy as np
from scipy.interpolate import make_interp_spline
from scipy import special
import seaborn as sns
import matplotlib.pyplot as plt


class Table():
    """Normal distribution analisis table. Solve task from university."""

    def __init__(self, data, _min=None, _max=None, strange_num=None)->None:
        if strange_num == None:
            strange_num = int(len(data)/10)
        self.data = data
        self.step = int((_max - _min) / strange_num)
        self.m = np.mean(data)
        self.D = np.var(data)
        self.sig = np.std(data)
        self._max = _max
        self._min = _min
        self.strange_num = strange_num
        self.cols = []

        col1 = [i for i in range(_min, _max, self.step)] 
        if col1[-1] < max(data):
            col1.append(col1[-1]+self.step)
        self.col1 = np.array(col1)
        self.cols.append(self.col1)

        col2 = []
        for i in range(len(col1)-1):
            col2.append(np.average([col1[i], col1[i+1]]))
        self.col2 = np.array(col2)
        self.cols.append(self.col2)
        
        col3 = []
        for i in range(len(col1)-1):
            a = col1[i]
            b = col1[i+1]
            col3.append(len([j for j in self.data if a < j <= b]))
        self.col3 = np.array(col3)
        self.cols.append(self.col3)
  
        col4 = [i/len(self.data) for i in col3]
        self.col4 = np.array(col4)
        self.cols.append(self.col4)

        self.col5 = np.array([col2[i] * col4[i] for i in range(len(col2))])
        self.cols.append(self.col5)

        col6 = [col2[i]**2 * col4[i] for i in range(len(col2))]
        self.col6 = np.array(col6)
        self.cols.append(self.col6)

        col7 = [(i-self.m)/self.sig for i in col1]
        self.col7 = np.array(col7)
        self.cols.append(self.col7)
        
        f = lambda x: special.erf(x/np.sqrt(2))/2
        col8 = [f(x) for x in col7]
        self.col8 = np.random.normal(col8)
        self.cols.append(self.col8)

        col9 = [np.absolute(col8[i]-col8[i+1]) for i in range(len(col8)-1)]
        self.col9 = np.array(col9)
        self.cols.append(self.col9)

        hi_col = [(col9[i] - col4[i])**2/col4[i] for i in range(len(col9))]
        self.hi = len(data) * sum(hi_col)
        self.hi_f = f"{len(data)}*{sum(hi_col)}"

    
    def toJS(self):
        heads = [
                 "xi",
                 "xti",
                 "nbi",
                 "pbieq",
                 "xtipbi",
                 "xti2pbi",
                 "xi-m",
                 "F",
                 "pi"
        ]
        s = ""
        for i, v in enumerate(self.cols):
          # Create colomn like `let col_i = {head: ..., data: ..., sum: ...}`
            s += \
'''let col{} = {{
    head: "{}",
    data: [{}],
    sum: {}
}}
            
'''.format(i+1, heads[i], ', '.join([str(i) for i in v]), sum(v))

        s += \
'''let info = {{
    H_f: "({}-{})/{}",
    H: {},
    max: {},
    min: {},
    sig: {},
    D: {},
    m: {},
    avr: {},
    hi: {}
}}
'''.format(self._max, self._min, 
                      self.strange_num, self.step, 
                      max(self.data), min(self.data),
                      self.sig, self.D,
                      sum(self.col5), self.m,
                      self.hi)
        return s

    def plot(self)-> None:
        plt.style.use('seaborn-whitegrid')

        # Create the figure and axes objects
        fig, ax = plt.subplots(1, figsize=(8, 6))
        fig.suptitle('Example Of Plot With Major and Minor Grid Lines')

        # Plot the data
        # Not smooth
        ax.plot(self.col2, self.col4, marker='o')

        # Smooth
        x = self.col2
        y = self.col9
        xnew = np.linspace(x.min(), x.max(), 200) 
        spl = make_interp_spline(x, y, k=3)
        y_smooth = spl(xnew)
        ax.plot(xnew, y_smooth)

        # Show the major grid lines with dark grey lines
        plt.grid(b=True, which='major', color='#666666', linestyle='-')

        # Show the minor grid lines with very faint and almost transparent grey lines
        plt.minorticks_on()
        plt.grid(b=True, which='minor', color='#999999', linestyle='-', alpha=0.7)

        plt.show()

    def __str__(self)->str:
        s = ""
        s += f"Step: {self.step}\n"
        s += f"m: {self.m}\n"
        s += f"D: {self.D}\n"
        s += f"Sig: {self.sig}\n"
        for i, v in enumerate(self.cols):
            s += f"Col {i+1}\n{v}\n∑={sum(v)}\n\n"
        return s

d = "223 228 220 196 188 193 195 199 236 205 230 178 203 172 213 219 202 159 197 187 259 181 207 196 197 160 190 180 202 191 194 164 216 201 231 181 248 208 188 198 195 237 227 204 165 201 204 226 284 189 217 177 203 229 170 201 181 193 170 201 235 225 186 219 191 212 183 219 226 171 234 185 201 255 193 181 220 209 186 188 205 251 190 196 245 175 152 297 176 229 258 233 191 200 190 211 185 196 172 215"
data = [int(i) for i in d.split(" ")]

config = input(f"Max is {max(data)}, min is {min(data)}. Enter your numbers > ")
digits = [int(i) for i in config.strip().split(' ')]
t = Table(data, _min=min(digits), _max=max(digits))
print(t.toJS())
t.plot()